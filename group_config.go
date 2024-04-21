package main

import (
	"flag"
	"regexp"
)

type GroupConfig struct {
	name 							string
	allowedSender 		*regexp.Regexp
	allowedRecipients *regexp.Regexp
	command           string
	remotes 					[]*Remote
}

type GroupConfigStr struct {
	name 							string
	allowedSenderStr *string
	allowedRecipStr  *string
	command          *string
	remotesStr  		 *string
}

func (g *GroupConfig) senderAllowed(addr string) bool {
	if g.allowedSender == nil {
		// Any sender is permitted
		return true
	}

	if g.allowedSender.MatchString(addr) {
		// Permitted by regex
		return true
	}
	
	return false
}

func (g *GroupConfig) recipientAllowed(addr string) bool {
	if g.allowedRecipients == nil {
		// Any recipient is permitted
		return true
	}

	if g.allowedRecipients.MatchString(addr) {
		// Permitted by regex
		return true
	}
	
	return false
}

func (g *GroupConfig) filterRecipients(recipients []string) (result []string) {
	for _, recip := range recipients {
		if g.recipientAllowed(recip) {
			result = append(result, recip)
		}
	}

	return
}

func AppendGroupFlags(flagset *flag.FlagSet, group string) *GroupConfigStr {
	config := &GroupConfigStr{}
	prefix := "g_" + group

	if group != "" {
		config.allowedSenderStr = flagset.String(prefix + "_allowed_sender", "", "Regular expression for valid FROM EMail addresses for group: " + group)
		config.allowedRecipStr  = flagset.String(prefix + "_allowed_recipients", "", "Regular expression for valid TO EMail addresses for group: " + group)
		config.command          = flagset.String(prefix + "_command", "", "Path to pipe command for group: " + group)
		config.remotesStr       = flagset.String(prefix + "_remotes", "", "Outgoing SMTP servers for group: " + group)
	} else {
		config.allowedSenderStr = flagset.String("allowed_sender", "", "Regular expression for valid FROM EMail addresses")
		config.allowedRecipStr  = flagset.String("allowed_recipients", "", "Regular expression for valid TO EMail addresses")
		config.command          = flagset.String("command", "", "Path to pipe command")
		config.remotesStr       = flagset.String("remotes", "", "Outgoing SMTP servers")
	}

	return config
}
