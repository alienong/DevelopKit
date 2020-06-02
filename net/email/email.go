/**
 * @Author: alienongwlx@gmail.com
 * @Description: Get Email Content By IMap Server
 * @Version: 1.0.0
 * @Date: 2020/4/15 19:50
 */

package email

import (
	"errors"
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"io"
	"io/ioutil"
	"time"
)

/**
@description: MailService Struct
@attribute Client: IMap Client
@attribute Account: Email Account
@attribute Password: Email Password
@attribute IsLogin: Login State
*/
type MailService struct {
	Client   *client.Client
	Account  string
	Password string
	IsLogin  bool
}

/**
@description: Mail Content Struct
@attribute Date: Email Date
@attribute From: Email Sender
@attribute To: Email Receiver
@attribute Subject: Email Subject
@attribute Content: Email Content
@attribute Attach: Email Attach List
*/
type MailContent struct {
	Date    *time.Time
	From    []*mail.Address
	To      []*mail.Address
	Subject string
	Content string
	Attach  []string
}

/**
@description: NewMailService
@param host: IMap Server
@param account: Email Account
@param password: Email Password
@return: MailService
*/
func NewMailService(host, account, password string) *MailService {
	c, err := client.DialTLS(host, nil)
	if err != nil {
		return nil
	}
	return &MailService{c, account, password, false}
}

/**
@description: Login The Email
@param : None
@return: error or nil
*/
func (ms *MailService) Login() error {
	err := ms.Client.Login(ms.Account, ms.Password)
	if err != nil {
		return err
	}
	ms.IsLogin = true
	return nil
}

/**
@description: Logout The Email
@param : None
@return: error or nil
*/
func (ms *MailService) Logout() error {
	if ms.IsLogin == true {
		ms.IsLogin = false
		return ms.Logout()
	}
	return nil
}

/**
@description: Get MailBox List
@param : None
@return: MailBox List And Error
*/
func (ms *MailService) FetchMailboxInfo() ([]string, error) {
	if !ms.IsLogin {
		err := ms.Login()
		if err != nil {
			return nil, err
		}
	}
	mis := make([]string, 0)
	done := make(chan error, 1)
	mailBoxes := make(chan *imap.MailboxInfo, 10)
	go func() {
		done <- ms.Client.List("", "*", mailBoxes)
	}()
	for m := range mailBoxes {
		mis = append(mis, m.Name)
	}
	if err := <-done; err != nil {
		close(done)
		return nil, err
	}
	return mis, nil
}

/**
@description: Get MailContent List
@param last: Last Count
@param boxName: Mail Box Name
@return: MailContent List
*/
func (ms *MailService) FetchMailContent(last uint32, boxName string) ([]*MailContent, error) {
	mcs := make([]*MailContent, 0)
	mbox, err := ms.Client.Select(boxName, false)
	if err != nil {
		return mcs, err
	}
	if mbox.Messages == 0 {
		return mcs, errors.New(fmt.Sprintf("%s has no mails", boxName))
	}
	start := uint32(0)
	if mbox.Messages > last {
		start = mbox.Messages - last
	}
	seqSet := new(imap.SeqSet)
	seqSet.AddRange(start, mbox.Messages)
	// Get the whole message body
	var section imap.BodySectionName
	items := []imap.FetchItem{section.FetchItem()}
	messages := make(chan *imap.Message, 10)
	go func() {
		if err := ms.Client.Fetch(seqSet, items, messages); err != nil {
			fmt.Println(err)
		}
	}()
	for msg := range messages {
		if msg == nil {
			return mcs, errors.New(fmt.Sprintf("%s server didn't returned message", boxName))
		}
		r := msg.GetBody(&section)
		if r == nil {
			return mcs, errors.New(fmt.Sprintf("%s server didn't returned message body", boxName))
		}
		// Create a new mail reader
		mr, err := mail.CreateReader(r)
		if err != nil {
			return mcs, err
		}
		// Get some info about the message
		mc := new(MailContent)
		header := mr.Header
		if date, err := header.Date(); err == nil {
			mc.Date = &date
		} else {
			continue
		}
		if from, err := header.AddressList("From"); err == nil {
			mc.From = from
		} else {
			continue
		}
		if to, err := header.AddressList("To"); err == nil {
			mc.To = to
		} else {
			continue
		}
		if subject, err := header.Subject(); err == nil {
			mc.Subject = subject
		} else {
			continue
		}
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			} else if err != nil {
				break
			}
			switch h := p.Header.(type) {
			case *mail.InlineHeader:
				b, _ := ioutil.ReadAll(p.Body)
				mc.Content = mc.Content + string(b)
			case *mail.AttachmentHeader:
				filename, _ := h.Filename()
				mc.Attach = append(mc.Attach, filename)
			}
			mcs = append(mcs, mc)
		}

	}
	return mcs, nil
}
