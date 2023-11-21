package cluster

import "go-redis/interface/resp"

type ClusterDatabase struct {
}

func MakeClusterDatabase() *ClusterDatabase {

}

func (c *ClusterDatabase) Exec(Client resp.Connection, args [][]byte) resp.Reply {
	//TODO implement me
	panic("implement me")
}

func (c *ClusterDatabase) Close() {
	//TODO implement me
	panic("implement me")
}

func (c *ClusterDatabase) AfterClientClose(conn resp.Connection) {
	//TODO implement me
	panic("implement me")
}
