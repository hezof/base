package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func LookupPath(name string) string {
	// 1. 如果name存在
	if ExistsPath(name) {
		return name
	}
	// 2. 当前目录查找
	if path := filepath.Join(filepath.Dir(os.Args[0]), name); ExistsPath(path) {
		return path
	}
	// 3. 工作目录查找
	if cwd, err := os.Getwd(); err == nil {
		if path := filepath.Join(cwd, name); ExistsPath(path) {
			return path
		}
	}
	// 4. 系统PATH查找
	if path, err := exec.LookPath(name); err == nil {
		if ExistsPath(path) {
			return path
		}
	}
	// 5. 上述步骤都失败,返回初值碰运气!
	return name
}

func LocatePath(env, def string) (string, error) {
	var path string
	// 1. 环境变量配置绝对路径
	path = os.Getenv(env)
	if path != "" {
		if ExistsPath(path) {
			return path, nil
		} else {
			return path, fmt.Errorf("%v not found", path)
		}
	}
	// 2. 启动目录的默认配置
	path = filepath.Join(filepath.Dir(os.Args[0]), def)
	if ExistsPath(path) {
		return path, nil
	}
	// 3. 工作目录的默认配置
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path = filepath.Join(cwd, def)
	if ExistsPath(path) {
		return path, nil
	}
	return "", nil
}

func ExistsPath(path string) bool {
	stat, err := os.Stat(path)
	return stat != nil || os.IsExist(err)
}
