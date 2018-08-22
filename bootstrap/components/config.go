package components

import (
    "github.com/tidwall/gjson"
    "strings"
)

type ConfigFacade string

var Config *ConfigFacade

func init() {
    Config = new(ConfigFacade)
}

func (cfg *ConfigFacade) Get(key string) gjson.Result {
    key_path := strings.Split(key, ".")

    if !FileSystem.PathExists("./config/" + key_path[0] + ".json") {
        return gjson.Result{}
    }

    body, _ := FileSystem.ReadAll("./config/" + key_path[0] + ".json")
    new_key := strings.Join(key_path[1:], ".")

    return gjson.Get(string(body), new_key)
}

func (cfg *ConfigFacade) GetArray(key string) []gjson.Result {
    key_path := strings.Split(key, ".")

    if !FileSystem.PathExists("./config/" + key_path[0] + ".json") {
        return []gjson.Result{}
    }

    body, _ := FileSystem.ReadAll("./config/" + key_path[0] + ".json")
    new_key := strings.Join(key_path[1:], ".")

    return gjson.Get(string(body), new_key).Array()
}

// ConfigFacade有点小用，可已缓存所有配置，不过也可以不缓存
