package types

type Room struct {
	Clients   map[*Client]bool
	Broadcast chan []byte
}

func (r *Room) Run() {
	for msg := range r.Broadcast {
		for client := range r.Clients {
			select {
			case client.Send <- msg:
			default:
				close(client.Send)
				delete(r.Clients, client)
			}
		}
	}
}
