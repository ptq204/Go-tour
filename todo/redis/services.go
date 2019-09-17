package redis

import (
	"fmt"
	"github.com/bradfitz/slice"
	"github.com/gomodule/redigo/redis"
)

// TaskRet exported
type TaskRet struct {
	Key     string
	Content string
}

// RedisInstance struct
type RedisInstance struct {
	pool *redis.Pool
	conn redis.Conn
}

// CreateInstance exported
func CreateInstance() *RedisInstance {
	pool := NewPool()
	conn := pool.Get()
	r := RedisInstance{pool, conn}
	err := r.Ping()
	if err != nil {
		fmt.Println(err)
	}
	return &r
}

// Ping exported
func (r *RedisInstance) Ping() error {
	pong, err := r.conn.Do("PING")
	if err != nil {
		return err
	}
	s, err := redis.String(pong, err)
	if err != nil {
		return err
	}
	fmt.Printf("PING response = %s\n", s)
	return nil
}

// SaveTask exported
func (r *RedisInstance) SaveTask(key string, content string) error {
	_, err := r.conn.Do("SET", key, content)
	if err != nil {
		return err
	}
	return nil
}

// GetTaskList exported
func (r *RedisInstance) GetTaskList() ([]TaskRet, error) {
	var keyList []string
	var taskList []TaskRet
	keyList, err := redis.Strings(r.conn.Do("KEYS", "todo:*"))
	if err != nil {
		return nil, err
	}
	for _, key := range keyList {
		s, err := redis.String(r.conn.Do("GET", key))
		if err != nil {
			return nil, err
		}
		taskList = append(taskList, TaskRet{key, s})
	}
	slice.Sort(taskList, func(i, j int) bool {
		return taskList[i].Key < taskList[j].Key
	})
	return taskList, nil
}

// MarkTaskDone exported
func (r *RedisInstance) MarkTaskDone(id int, taskList []TaskRet) error {
	id--
	if id >= 0 && id < len(taskList) {
		task := taskList[id]
		fmt.Printf("You have completed the \"%s\" task\n", task.Content)
		_, err := r.conn.Do("DEL", task.Key)
		if err != nil {
			return err
		}
	}
	return nil
}

// Close connection
func (r *RedisInstance) Close() {
	r.conn.Close()
}
