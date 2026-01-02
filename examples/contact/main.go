package main

import (
	"context"
	"fmt"
	"log"

	"github.com/shuaidd/wecom-core"
	"github.com/shuaidd/wecom-core/config"
	"github.com/shuaidd/wecom-core/pkg/logger"
	"github.com/shuaidd/wecom-core/types/contact"
)

func main() {
	// 1. 创建企业微信客户端
	client, err := wecom.New(
		config.WithCorpID("your_corp_id"),
		config.WithCorpSecret("your_corp_secret"),
		config.WithLogger(logger.NewStdLogger()),
		config.WithRetry(3),
	)
	if err != nil {
		log.Fatalf("Failed to create wecom client: %v", err)
	}

	ctx := context.Background()

	// 2. 创建成员
	fmt.Println("=== 创建成员 ===")
	createResp, err := client.Contact.CreateUser(ctx, &contact.CreateUserRequest{
		UserID:     "zhangsan",
		Name:       "张三",
		Mobile:     "13800000000",
		Department: []int{1},
		Email:      "zhangsan@example.com",
		Position:   "产品经理",
	})
	if err != nil {
		log.Printf("Failed to create user: %v", err)
	} else {
		fmt.Printf("成员创建成功: %+v\n", createResp)
	}

	// 3. 读取成员
	fmt.Println("\n=== 读取成员 ===")
	user, err := client.Contact.GetUser(ctx, "zhangsan")
	if err != nil {
		log.Printf("Failed to get user: %v", err)
	} else {
		fmt.Printf("成员信息: UserID=%s, Name=%s, Mobile=%s, Email=%s\n",
			user.UserID, user.Name, user.Mobile, user.Email)
	}

	// 4. 更新成员
	fmt.Println("\n=== 更新成员 ===")
	err = client.Contact.UpdateUser(ctx, &contact.UpdateUserRequest{
		UserID:   "zhangsan",
		Position: "高级产品经理",
	})
	if err != nil {
		log.Printf("Failed to update user: %v", err)
	} else {
		fmt.Println("成员更新成功")
	}

	// 5. 获取部门成员列表
	fmt.Println("\n=== 获取部门成员列表 ===")
	users, err := client.Contact.ListUsers(ctx, 1, false)
	if err != nil {
		log.Printf("Failed to list users: %v", err)
	} else {
		fmt.Printf("部门成员数量: %d\n", len(users))
		for _, u := range users {
			fmt.Printf("  - UserID=%s, Name=%s\n", u.UserID, u.Name)
		}
	}

	// 6. 创建部门
	fmt.Println("\n=== 创建部门 ===")
	deptID, err := client.Contact.CreateDepartment(ctx, &contact.CreateDepartmentRequest{
		Name:     "研发部",
		ParentID: 1,
		Order:    1,
	})
	if err != nil {
		log.Printf("Failed to create department: %v", err)
	} else {
		fmt.Printf("部门创建成功, ID: %d\n", deptID)
	}

	// 7. 获取部门列表
	fmt.Println("\n=== 获取部门列表 ===")
	departments, err := client.Contact.ListDepartments(ctx, 1)
	if err != nil {
		log.Printf("Failed to list departments: %v", err)
	} else {
		fmt.Printf("部门数量: %d\n", len(departments))
		for _, dept := range departments {
			fmt.Printf("  - ID=%d, Name=%s, ParentID=%d\n",
				dept.ID, dept.Name, dept.ParentID)
		}
	}

	// 8. 删除成员
	fmt.Println("\n=== 删除成员 ===")
	err = client.Contact.DeleteUser(ctx, "zhangsan")
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
	} else {
		fmt.Println("成员删除成功")
	}

	fmt.Println("\n=== 示例完成 ===")
}
