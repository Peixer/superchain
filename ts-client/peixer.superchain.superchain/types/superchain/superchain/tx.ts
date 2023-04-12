/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "peixer.superchain.superchain";

export interface MsgCreatePythonCode {
  creator: string;
  uRI: string;
}

export interface MsgCreatePythonCodeResponse {
  id: number;
}

export interface MsgUpdatePythonCode {
  creator: string;
  id: number;
  uRI: string;
}

export interface MsgUpdatePythonCodeResponse {
}

export interface MsgDeletePythonCode {
  creator: string;
  id: number;
}

export interface MsgDeletePythonCodeResponse {
}

export interface MsgRunCode {
  creator: string;
  id: number;
}

export interface MsgRunCodeResponse {
}

function createBaseMsgCreatePythonCode(): MsgCreatePythonCode {
  return { creator: "", uRI: "" };
}

export const MsgCreatePythonCode = {
  encode(message: MsgCreatePythonCode, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.uRI !== "") {
      writer.uint32(18).string(message.uRI);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreatePythonCode {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreatePythonCode();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.uRI = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreatePythonCode {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      uRI: isSet(object.uRI) ? String(object.uRI) : "",
    };
  },

  toJSON(message: MsgCreatePythonCode): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.uRI !== undefined && (obj.uRI = message.uRI);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreatePythonCode>, I>>(object: I): MsgCreatePythonCode {
    const message = createBaseMsgCreatePythonCode();
    message.creator = object.creator ?? "";
    message.uRI = object.uRI ?? "";
    return message;
  },
};

function createBaseMsgCreatePythonCodeResponse(): MsgCreatePythonCodeResponse {
  return { id: 0 };
}

export const MsgCreatePythonCodeResponse = {
  encode(message: MsgCreatePythonCodeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreatePythonCodeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreatePythonCodeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreatePythonCodeResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgCreatePythonCodeResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreatePythonCodeResponse>, I>>(object: I): MsgCreatePythonCodeResponse {
    const message = createBaseMsgCreatePythonCodeResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgUpdatePythonCode(): MsgUpdatePythonCode {
  return { creator: "", id: 0, uRI: "" };
}

export const MsgUpdatePythonCode = {
  encode(message: MsgUpdatePythonCode, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.uRI !== "") {
      writer.uint32(26).string(message.uRI);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdatePythonCode {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdatePythonCode();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.uRI = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdatePythonCode {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
      uRI: isSet(object.uRI) ? String(object.uRI) : "",
    };
  },

  toJSON(message: MsgUpdatePythonCode): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.uRI !== undefined && (obj.uRI = message.uRI);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdatePythonCode>, I>>(object: I): MsgUpdatePythonCode {
    const message = createBaseMsgUpdatePythonCode();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    message.uRI = object.uRI ?? "";
    return message;
  },
};

function createBaseMsgUpdatePythonCodeResponse(): MsgUpdatePythonCodeResponse {
  return {};
}

export const MsgUpdatePythonCodeResponse = {
  encode(_: MsgUpdatePythonCodeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdatePythonCodeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdatePythonCodeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdatePythonCodeResponse {
    return {};
  },

  toJSON(_: MsgUpdatePythonCodeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdatePythonCodeResponse>, I>>(_: I): MsgUpdatePythonCodeResponse {
    const message = createBaseMsgUpdatePythonCodeResponse();
    return message;
  },
};

function createBaseMsgDeletePythonCode(): MsgDeletePythonCode {
  return { creator: "", id: 0 };
}

export const MsgDeletePythonCode = {
  encode(message: MsgDeletePythonCode, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeletePythonCode {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeletePythonCode();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeletePythonCode {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgDeletePythonCode): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeletePythonCode>, I>>(object: I): MsgDeletePythonCode {
    const message = createBaseMsgDeletePythonCode();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgDeletePythonCodeResponse(): MsgDeletePythonCodeResponse {
  return {};
}

export const MsgDeletePythonCodeResponse = {
  encode(_: MsgDeletePythonCodeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeletePythonCodeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeletePythonCodeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeletePythonCodeResponse {
    return {};
  },

  toJSON(_: MsgDeletePythonCodeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeletePythonCodeResponse>, I>>(_: I): MsgDeletePythonCodeResponse {
    const message = createBaseMsgDeletePythonCodeResponse();
    return message;
  },
};

function createBaseMsgRunCode(): MsgRunCode {
  return { creator: "", id: 0 };
}

export const MsgRunCode = {
  encode(message: MsgRunCode, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).int32(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRunCode {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRunCode();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRunCode {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgRunCode): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRunCode>, I>>(object: I): MsgRunCode {
    const message = createBaseMsgRunCode();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgRunCodeResponse(): MsgRunCodeResponse {
  return {};
}

export const MsgRunCodeResponse = {
  encode(_: MsgRunCodeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRunCodeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRunCodeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgRunCodeResponse {
    return {};
  },

  toJSON(_: MsgRunCodeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRunCodeResponse>, I>>(_: I): MsgRunCodeResponse {
    const message = createBaseMsgRunCodeResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreatePythonCode(request: MsgCreatePythonCode): Promise<MsgCreatePythonCodeResponse>;
  UpdatePythonCode(request: MsgUpdatePythonCode): Promise<MsgUpdatePythonCodeResponse>;
  DeletePythonCode(request: MsgDeletePythonCode): Promise<MsgDeletePythonCodeResponse>;
  RunCode(request: MsgRunCode): Promise<MsgRunCodeResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreatePythonCode = this.CreatePythonCode.bind(this);
    this.UpdatePythonCode = this.UpdatePythonCode.bind(this);
    this.DeletePythonCode = this.DeletePythonCode.bind(this);
    this.RunCode = this.RunCode.bind(this);
  }
  CreatePythonCode(request: MsgCreatePythonCode): Promise<MsgCreatePythonCodeResponse> {
    const data = MsgCreatePythonCode.encode(request).finish();
    const promise = this.rpc.request("peixer.superchain.superchain.Msg", "CreatePythonCode", data);
    return promise.then((data) => MsgCreatePythonCodeResponse.decode(new _m0.Reader(data)));
  }

  UpdatePythonCode(request: MsgUpdatePythonCode): Promise<MsgUpdatePythonCodeResponse> {
    const data = MsgUpdatePythonCode.encode(request).finish();
    const promise = this.rpc.request("peixer.superchain.superchain.Msg", "UpdatePythonCode", data);
    return promise.then((data) => MsgUpdatePythonCodeResponse.decode(new _m0.Reader(data)));
  }

  DeletePythonCode(request: MsgDeletePythonCode): Promise<MsgDeletePythonCodeResponse> {
    const data = MsgDeletePythonCode.encode(request).finish();
    const promise = this.rpc.request("peixer.superchain.superchain.Msg", "DeletePythonCode", data);
    return promise.then((data) => MsgDeletePythonCodeResponse.decode(new _m0.Reader(data)));
  }

  RunCode(request: MsgRunCode): Promise<MsgRunCodeResponse> {
    const data = MsgRunCode.encode(request).finish();
    const promise = this.rpc.request("peixer.superchain.superchain.Msg", "RunCode", data);
    return promise.then((data) => MsgRunCodeResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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
