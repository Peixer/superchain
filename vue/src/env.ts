const apiURL = import.meta.env.VITE_API_COSMOS ?? "http://164.90.168.110:1317";
const rpcURL = import.meta.env.VITE_WS_TENDERMINT ?? "http://164.90.168.110:26657";
const prefix = import.meta.env.VITE_ADDRESS_PREFIX ?? "cosmos";

export const env = {
  apiURL,
  rpcURL,
  prefix,
};
