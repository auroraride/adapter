// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-22
// Based on bmservd by liasica, magicrolan@qq.com.

package ent

import (
    "context"
    "entgo.io/ent/entc/integration/ent/migrate"
    _ "github.com/lib/pq"
    "log"
)

func New(dsn string, debug bool) *Client {
    ctx := context.Background()
    client, err := Open("postgres", dsn)
    if err != nil {
        log.Fatalf("数据库打开失败: %v", err)
    }

    if debug {
        client = client.Debug()
    }

    if err = client.Schema.Create(
        ctx,
        migrate.WithDropIndex(true),
        migrate.WithDropColumn(true),
        migrate.WithForeignKeys(false),
        // schema.WithApplyHook(func(next schema.Applier) schema.Applier {
        //     return schema.ApplyFunc(func(ctx context.Context, conn dialect.ExecQuerier, plan *sm.Plan) error {
        //         isCreateBattery := false
        //         isCreateHeartbeat := false
        //         for _, c := range plan.Changes {
        //             if c.Comment == `create "battery" table` {
        //                 isCreateBattery = true
        //                 break
        //             }
        //             if c.Comment == `create "heartbeat" table` {
        //                 isCreateHeartbeat = true
        //                 break
        //             }
        //         }
        //         err := next.Apply(ctx, conn, plan)
        //         if err == nil {
        //             if isCreateBattery {
        //                 return conn.Exec(ctx, `CREATE INDEX ON battery((cabinet->>'serial'))`, []any{}, nil)
        //             }
        //             if isCreateHeartbeat {
        //                 return conn.Exec(ctx, `CREATE INDEX ON heartbeat((cabinet->>'serial'))`, []any{}, nil)
        //             }
        //         }
        //         return err
        //     })
        // }),
    ); err != nil {
        log.Fatalf("数据库迁移失败: %v", err)
    }

    return client
}
