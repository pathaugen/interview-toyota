
package main

import (
)


// Stock Struct
type Stock struct {
  DomainName        string

  Title             string
  ShortDescription  string
  LongDescription   string
  Keywords          string // CSV

  GoogleCSE         string // CSE key
  WebmasterTools    string // google-site-verification key - https://support.google.com/webmasters/answer/35179 - <meta name="google-site-verification" content="String_we_ask_for">
  Analytics         string // key

  // Contact? -> Link to User, or have UsersInWebSiteDomains: User, WebsiteDomain, Role (Admin, Staff, Web Content, Web Templates)

  CreatedBy         string // TODO: LINK to Users
}
