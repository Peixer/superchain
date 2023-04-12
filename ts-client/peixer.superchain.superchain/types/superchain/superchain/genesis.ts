/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Params } from "./params";
import { PythonCode } from "./python_code";

export const protobufPackage = "peixer.superchain.superchain";

/** GenesisState defines the superchain module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  pythonCodeList: PythonCode[];
  pythonCodeCount: number;
}

function createBaseGenesisState(): GenesisState {
  return { params: undefined, pythonCodeList: [], pythonCodeCount: 0 };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.pythonCodeList) {
      PythonCode.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.pythonCodeCount !== 0) {
      writer.uint32(24).uint64(message.pythonCodeCount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.pythonCodeList.push(PythonCode.decode(reader, reader.uint32()));
          break;
        case 3:
          message.pythonCodeCount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      pythonCodeList: Array.isArray(object?.pythonCodeList)
        ? object.pythonCodeList.map((e: any) => PythonCode.fromJSON(e))
        : [],
      pythonCodeCount: isSet(object.pythonCodeCount) ? Number(object.pythonCodeCount) : 0,
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.pythonCodeList) {
      obj.pythonCodeList = message.pythonCodeList.map((e) => e ? PythonCode.toJSON(e) : undefined);
    } else {
      obj.pythonCodeList = [];
    }
    message.pythonCodeCount !== undefined && (obj.pythonCodeCount = Math.round(message.pythonCodeCount));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.pythonCodeList = object.pythonCodeList?.map((e) => PythonCode.fromPartial(e)) || [];
    message.pythonCodeCount = object.pythonCodeCount ?? 0;
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
