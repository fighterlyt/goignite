package gigocql

import (
	"context"
	"strings"

	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/gocql/gocql"
)

const (
	TopicSession = "topic:gocql:session"
)

func NewSession(ctx context.Context, o *Options) (session *gocql.Session, err error) {

	l := gilog.FromContext(ctx)

	cluster := gocql.NewCluster(o.Hosts...)

	if o.Port > 0 {
		cluster.Port = o.Port
	}

	if o.CQLVersion != "" {
		cluster.CQLVersion = o.CQLVersion
	}

	if o.ProtoVersion > 0 {
		cluster.ProtoVersion = o.ProtoVersion
	}

	if o.Timeout > 0 {
		cluster.Timeout = o.Timeout
	}

	if o.ConnectTimeout > 0 {
		cluster.ConnectTimeout = o.ConnectTimeout
	}

	if o.Keyspace != "" {
		cluster.Keyspace = o.Keyspace
	}

	if o.NumConns > 0 {
		cluster.NumConns = o.NumConns
	}

	if o.SocketKeepalive > 0 {
		cluster.SocketKeepalive = o.SocketKeepalive
	}

	if o.MaxPreparedStmts > 0 {
		cluster.MaxPreparedStmts = o.MaxPreparedStmts
	}

	if o.MaxRoutingKeyInfo > 0 {
		cluster.MaxRoutingKeyInfo = o.MaxRoutingKeyInfo
	}

	if o.PageSize > 0 {
		cluster.PageSize = o.PageSize
	}

	cluster.DefaultTimestamp = o.DefaultTimestamp

	if o.ReconnectInterval > 0 {
		cluster.ReconnectInterval = o.ReconnectInterval
	}

	if o.MaxWaitSchemaAgreement > 0 {
		cluster.MaxWaitSchemaAgreement = o.MaxWaitSchemaAgreement
	}

	cluster.DisableInitialHostLookup = o.DisableInitialHostLookup

	if o.WriteCoalesceWaitTime > 0 {
		cluster.WriteCoalesceWaitTime = o.WriteCoalesceWaitTime
	}

	if o.Username != "" || o.Password != "" {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: o.Username,
			Password: o.Password,
		}
	}

	if o.Consistency != "" {
		cluster.Consistency = gocql.ParseConsistency(o.Consistency)
	}

	session, err = cluster.CreateSession()

	if err != nil {
		return nil, err
	}

	gieventbus.Publish(TopicSession, session)

	l.Infof("Connected to Cassandra server: %v", strings.Join(o.Hosts, ","))

	return session, err
}

func NewDefaultSession(ctx context.Context) (*gocql.Session, error) {

	l := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewSession(ctx, o)
}
