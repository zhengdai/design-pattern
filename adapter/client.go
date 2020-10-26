package adapter

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Clinet inserts Lightning connector into computer.")
	com.insertIntoLightningPort()
}
