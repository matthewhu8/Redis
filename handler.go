package main

func ping(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: "string", str: "PONG"}
	}
	return Value{typ: "string", str: args[0].bulk}
}

func set(args []Value) Value {
	return Value{typ: "string", str: args[0].bulk}
}

func get(args []Value) Value {
	return Value{typ: "string", str: args[0].bulk}
}

var Handlers = map[string]func([]Value) Value{
	"PING": ping,
	"GET":  get,
	"SET":  set,
}
