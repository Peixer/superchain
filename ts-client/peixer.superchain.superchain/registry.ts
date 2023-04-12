import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreatePythonCode } from "./types/superchain/superchain/tx";
import { MsgDeletePythonCode } from "./types/superchain/superchain/tx";
import { MsgRunCode } from "./types/superchain/superchain/tx";
import { MsgUpdatePythonCode } from "./types/superchain/superchain/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/peixer.superchain.superchain.MsgCreatePythonCode", MsgCreatePythonCode],
    ["/peixer.superchain.superchain.MsgDeletePythonCode", MsgDeletePythonCode],
    ["/peixer.superchain.superchain.MsgRunCode", MsgRunCode],
    ["/peixer.superchain.superchain.MsgUpdatePythonCode", MsgUpdatePythonCode],
    
];

export { msgTypes }