package core

import (
	"fmt"
	"github.com/hezof/log"
	"time"
)

/**************************************************
 * 通用配置
 **************************************************/

func ConfigBool(path string) (bool, error) {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		return false, fmt.Errorf("config unknown %v", path)
	}
	ret, err := AsBool(val)
	if err != nil {
		return false, fmt.Errorf("config invalid %v, %v", path, err)
	}
	return ret, nil
}

func ConfigString(path string) (string, error) {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		return "", fmt.Errorf("config unknown %v", path)
	}
	ret, err := AsString(val)
	if err != nil {
		return "", fmt.Errorf("config invalid %v, %v", path, err)
	}
	return ret, nil
}

func ConfigTime(path string) (time.Time, error) {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		return ZeroTime, fmt.Errorf("config unknown %v", path)
	}
	ret, err := AsTime(val)
	if err != nil {
		return ZeroTime, fmt.Errorf("config invalid %v, %v", path, err)
	}
	return ret, nil
}

func ConfigDuration(path string) (time.Duration, error) {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		return 0, fmt.Errorf("config unknown %v", path)
	}
	ret, err := AsDuration(val)
	if err != nil {
		return 0, fmt.Errorf("config invalid %v, %v", path, err)
	}
	return ret, nil
}

func ConfigInt(path string) (int, error) {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		return 0, fmt.Errorf("config unknown %v", path)
	}
	ret, err := AsInt(val)
	if err != nil {
		return 0, fmt.Errorf("config invalid %v, %v", path, err)
	}
	return ret, nil
}

func ConfigInt8(path string) (int8, error) {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		return 0, fmt.Errorf("config unknown %v", path)
	}
	ret, err := AsInt8(val)
	if err != nil {
		return 0, fmt.Errorf("config invalid %v, %v", path, err)
	}
	return ret, nil
}

func ConfigInt16(path string) (int16, error) {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		return 0, fmt.Errorf("config unknown %v", path)
	}
	ret, err := AsInt16(val)
	if err != nil {
		return 0, fmt.Errorf("config invalid %v, %v", path, err)
	}
	return ret, nil
}

func ConfigInt32(path string) (int32, error) {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		return 0, fmt.Errorf("config unknown %v", path)
	}
	ret, err := AsInt32(val)
	if err != nil {
		return 0, fmt.Errorf("config invalid %v, %v", path, err)
	}
	return ret, nil
}

func ConfigInt64(path string) (int64, error) {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		return 0, fmt.Errorf("config unknown %v", path)
	}
	ret, err := AsInt64(val)
	if err != nil {
		return 0, fmt.Errorf("config invalid %v, %v", path, err)
	}
	return ret, nil
}

func ConfigUint(path string) (uint, error) {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		return 0, fmt.Errorf("config unknown %v", path)
	}
	ret, err := AsUint(val)
	if err != nil {
		return 0, fmt.Errorf("config invalid %v, %v", path, err)
	}
	return ret, nil
}

func ConfigUint8(path string) (uint8, error) {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		return 0, fmt.Errorf("config unknown %v", path)
	}
	ret, err := AsUint8(val)
	if err != nil {
		return 0, fmt.Errorf("config invalid %v, %v", path, err)
	}
	return ret, nil
}

func ConfigUint16(path string) (uint16, error) {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		return 0, fmt.Errorf("config unknown %v", path)
	}
	ret, err := AsUint16(val)
	if err != nil {
		return 0, fmt.Errorf("config invalid %v, %v", path, err)
	}
	return ret, nil
}

func ConfigUint32(path string) (uint32, error) {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		return 0, fmt.Errorf("config unknown %v", path)
	}
	ret, err := AsUint32(val)
	if err != nil {
		return 0, fmt.Errorf("config invalid %v, %v", path, err)
	}
	return ret, nil
}

func ConfigUint64(path string) (uint64, error) {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		return 0, fmt.Errorf("config unknown %v", path)
	}
	ret, err := AsUint64(val)
	if err != nil {
		return 0, fmt.Errorf("config invalid %v, %v", path, err)
	}
	return ret, nil
}

func ConfigFloat32(path string) (float32, error) {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		return 0, fmt.Errorf("config unknown %v", path)
	}
	ret, err := AsFloat32(val)
	if err != nil {
		return 0, fmt.Errorf("config invalid %v, %v", path, err)
	}
	return ret, nil
}

func ConfigFloat64(path string) (float64, error) {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		return 0, fmt.Errorf("config unknown %v", path)
	}
	ret, err := AsFloat64(val)
	if err != nil {
		return 0, fmt.Errorf("config invalid %v, %v", path, err)
	}
	return ret, nil
}

/**************************************************
 * 断言配置
 **************************************************/

func MustConfigBool(path string) bool {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		panic(fmt.Errorf("config unknown %v", path))
	}
	ret, err := AsBool(val)
	if err != nil {
		panic(fmt.Errorf("config invalid %v, %v", path, err))
	}
	return ret
}

func MustConfigString(path string) string {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		panic(fmt.Errorf("config unknown %v", path))
	}
	ret, err := AsString(val)
	if err != nil {
		panic(fmt.Errorf("config invalid %v, %v", path, err))
	}
	return ret
}

func MustConfigTime(path string) time.Time {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		panic(fmt.Errorf("config unknown %v", path))
	}
	ret, err := AsTime(val)
	if err != nil {
		panic(fmt.Errorf("config invalid %v, %v", path, err))
	}
	return ret
}

func MustConfigDuration(path string) time.Duration {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		panic(fmt.Errorf("config unknown %v", path))
	}
	ret, err := AsDuration(val)
	if err != nil {
		panic(fmt.Errorf("config invalid %v, %v", path, err))
	}
	return ret
}

func MustConfigInt(path string) int {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		panic(fmt.Errorf("config unknown %v", path))
	}
	ret, err := AsInt(val)
	if err != nil {
		panic(fmt.Errorf("config invalid %v, %v", path, err))
	}
	return ret
}

func MustConfigInt8(path string) int8 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		panic(fmt.Errorf("config unknown %v", path))
	}
	ret, err := AsInt8(val)
	if err != nil {
		panic(fmt.Errorf("config invalid %v, %v", path, err))
	}
	return ret
}

func MustConfigInt16(path string) int16 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		panic(fmt.Errorf("config unknown %v", path))
	}
	ret, err := AsInt16(val)
	if err != nil {
		panic(fmt.Errorf("config invalid %v, %v", path, err))
	}
	return ret
}

func MustConfigInt32(path string) int32 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		panic(fmt.Errorf("config unknown %v", path))
	}
	ret, err := AsInt32(val)
	if err != nil {
		panic(fmt.Errorf("config invalid %v, %v", path, err))
	}
	return ret
}

func MustConfigInt64(path string) int64 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		panic(fmt.Errorf("config unknown %v", path))
	}
	ret, err := AsInt64(val)
	if err != nil {
		panic(fmt.Errorf("config invalid %v, %v", path, err))
	}
	return ret
}

func MustConfigUint(path string) uint {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		panic(fmt.Errorf("config unknown %v", path))
	}
	ret, err := AsUint(val)
	if err != nil {
		panic(fmt.Errorf("config invalid %v, %v", path, err))
	}
	return ret
}

func MustConfigUint8(path string) uint8 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		panic(fmt.Errorf("config unknown %v", path))
	}
	ret, err := AsUint8(val)
	if err != nil {
		panic(fmt.Errorf("config invalid %v, %v", path, err))
	}
	return ret
}

func MustConfigUint16(path string) uint16 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		panic(fmt.Errorf("config unknown %v", path))
	}
	ret, err := AsUint16(val)
	if err != nil {
		panic(fmt.Errorf("config invalid %v, %v", path, err))
	}
	return ret
}

func MustConfigUint32(path string) uint32 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		panic(fmt.Errorf("config unknown %v", path))
	}
	ret, err := AsUint32(val)
	if err != nil {
		panic(fmt.Errorf("config invalid %v, %v", path, err))
	}
	return ret
}

func MustConfigUint64(path string) uint64 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		panic(fmt.Errorf("config unknown %v", path))
	}
	ret, err := AsUint64(val)
	if err != nil {
		panic(fmt.Errorf("config invalid %v, %v", path, err))
	}
	return ret
}

func MustConfigFloat32(path string) float32 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		panic(fmt.Errorf("config unknown %v", path))
	}
	ret, err := AsFloat32(val)
	if err != nil {
		panic(fmt.Errorf("config invalid %v, %v", path, err))
	}
	return ret
}

func MustConfigFloat64(path string) float64 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		panic(fmt.Errorf("config unknown %v", path))
	}
	ret, err := AsFloat64(val)
	if err != nil {
		panic(fmt.Errorf("config invalid %v, %v", path, err))
	}
	return ret
}

/**************************************************
 * 可选配置
 **************************************************/

func OptiConfigBool(path string, def bool) bool {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		log.Warn("config unknown %v", path)
		return def
	}
	ret, err := AsBool(val)
	if err != nil {
		log.Warn("config invalid %v, %v", path, err)
		return def
	}
	return ret
}

func OptiConfigString(path string, def string) string {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		log.Warn("config unknown %v", path)
		return def
	}
	ret, err := AsString(val)
	if err != nil {
		log.Warn("config invalid %v, %v", path, err)
		return def
	}
	return ret
}

func OptiConfigTime(path string, def time.Time) time.Time {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		log.Warn("config unknown %v", path)
		return def
	}
	ret, err := AsTime(val)
	if err != nil {
		log.Warn("config invalid %v, %v", path, err)
		return def
	}
	return ret
}

func OptiConfigDuration(path string, def time.Duration) time.Duration {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		log.Warn("config unknown %v", path)
		return def
	}
	ret, err := AsDuration(val)
	if err != nil {
		log.Warn("config invalid %v, %v", path, err)
		return def
	}
	return ret
}

func OptiConfigInt(path string, def int) int {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		log.Warn("config unknown %v", path)
		return def
	}
	ret, err := AsInt(val)
	if err != nil {
		log.Warn("config invalid %v, %v", path, err)
		return def
	}
	return ret
}

func OptiConfigInt8(path string, def int8) int8 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		log.Warn("config unknown %v", path)
		return def
	}
	ret, err := AsInt8(val)
	if err != nil {
		log.Warn("config invalid %v, %v", path, err)
		return def
	}
	return ret
}

func OptiConfigInt16(path string, def int16) int16 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		log.Warn("config unknown %v", path)
		return def
	}
	ret, err := AsInt16(val)
	if err != nil {
		log.Warn("config invalid %v, %v", path, err)
		return def
	}
	return ret
}

func OptiConfigInt32(path string, def int32) int32 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		log.Warn("config unknown %v", path)
		return def
	}
	ret, err := AsInt32(val)
	if err != nil {
		log.Warn("config invalid %v, %v", path, err)
		return def
	}
	return ret
}

func OptiConfigInt64(path string, def int64) int64 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		log.Warn("config unknown %v", path)
		return def
	}
	ret, err := AsInt64(val)
	if err != nil {
		log.Warn("config invalid %v, %v", path, err)
		return def
	}
	return ret
}

func OptiConfigUint(path string, def uint) uint {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		log.Warn("config unknown %v", path)
		return def
	}
	ret, err := AsUint(val)
	if err != nil {
		log.Warn("config invalid %v, %v", path, err)
		return def
	}
	return ret
}

func OptiConfigUint8(path string, def uint8) uint8 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		log.Warn("config unknown %v", path)
		return def
	}
	ret, err := AsUint8(val)
	if err != nil {
		log.Warn("config invalid %v, %v", path, err)
		return def
	}
	return ret
}

func OptiConfigUint16(path string, def uint16) uint16 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		log.Warn("config unknown %v", path)
		return def
	}
	ret, err := AsUint16(val)
	if err != nil {
		log.Warn("config invalid %v, %v", path, err)
		return def
	}
	return ret
}

func OptiConfigUint32(path string, def uint32) uint32 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		log.Warn("config unknown %v", path)
		return def
	}
	ret, err := AsUint32(val)
	if err != nil {
		log.Warn("config invalid %v, %v", path, err)
		return def
	}
	return ret
}

func OptiConfigUint64(path string, def uint64) uint64 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		log.Warn("config unknown %v", path)
		return def
	}
	ret, err := AsUint64(val)
	if err != nil {
		log.Warn("config invalid %v, %v", path, err)
		return def
	}
	return ret
}

func OptiConfigFloat32(path string, def float32) float32 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		log.Warn("config unknown %v", path)
		return def
	}
	ret, err := AsFloat32(val)
	if err != nil {
		log.Warn("config invalid %v, %v", path, err)
		return def
	}
	return ret
}

func OptiConfigFloat64(path string, def float64) float64 {
	val, ok := _configToml.GetFirst(path)
	if !ok {
		log.Warn("config unknown %v", path)
		return def
	}
	ret, err := AsFloat64(val)
	if err != nil {
		log.Warn("config invalid %v, %v", path, err)
		return def
	}
	return ret
}
