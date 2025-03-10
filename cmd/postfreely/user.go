/*
 * Copyright © 2020-2021 Musing Studio LLC.
 *
 * This file is part of WriteFreely.
 *
 * WriteFreely is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, included
 * in the LICENSE file in this source code package.
 */

package main

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/postfreely/postfreely"
)

var (
	cmdUser = cli.Command{
		Name:  "user",
		Usage: "user management tools",
		Subcommands: []*cli.Command{
			&cmdAddUser,
			&cmdDelUser,
			&cmdResetPass,
			// TODO: possibly add a user list command
		},
	}

	cmdAddUser = cli.Command{
		Name:    "create",
		Usage:   "Add new user",
		Aliases: []string{"a", "add"},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "admin",
				Value: false,
				Usage: "Create admin user",
			},
		},
		Action: addUserAction,
	}

	cmdDelUser = cli.Command{
		Name:    "delete",
		Usage:   "Delete user",
		Aliases: []string{"del", "d"},
		Action:  delUserAction,
	}

	cmdResetPass = cli.Command{
		Name:    "reset-pass",
		Usage:   "Reset user's password",
		Aliases: []string{"resetpass", "reset"},
		Action:  resetPassAction,
	}
)

func addUserAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return fmt.Errorf("No user passed. Example: writefreely user add [USER]:[PASSWORD]")
	}
	credentials := c.Args().Get(0)
	username, password, err := parseCredentials(credentials)
	if err != nil {
		return err
	}
	app := postfreely.NewApp(c.String("c"))
	return postfreely.CreateUser(app, username, password, c.Bool("admin"))
}

func delUserAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return fmt.Errorf("No user passed. Example: writefreely user delete [USER]")
	}
	username := c.Args().Get(0)
	app := postfreely.NewApp(c.String("c"))
	return postfreely.DoDeleteAccount(app, username)
}

func resetPassAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return fmt.Errorf("No user passed. Example: writefreely user reset-pass [USER]")
	}
	username := c.Args().Get(0)
	app := postfreely.NewApp(c.String("c"))
	return postfreely.ResetPassword(app, username)
}
