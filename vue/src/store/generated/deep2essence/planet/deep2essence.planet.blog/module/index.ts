// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateSentPost } from "./types/blog/tx";
import { MsgUpdatePost } from "./types/blog/tx";
import { MsgCreatePost } from "./types/blog/tx";
import { MsgCreateSentPost } from "./types/blog/tx";
import { MsgDeleteSentPost } from "./types/blog/tx";
import { MsgUpdateTimedoutPost } from "./types/blog/tx";
import { MsgDeletePost } from "./types/blog/tx";
import { MsgDeleteTimedoutPost } from "./types/blog/tx";
import { MsgCreateTimedoutPost } from "./types/blog/tx";


const types = [
  ["/deep2essence.planet.blog.MsgUpdateSentPost", MsgUpdateSentPost],
  ["/deep2essence.planet.blog.MsgUpdatePost", MsgUpdatePost],
  ["/deep2essence.planet.blog.MsgCreatePost", MsgCreatePost],
  ["/deep2essence.planet.blog.MsgCreateSentPost", MsgCreateSentPost],
  ["/deep2essence.planet.blog.MsgDeleteSentPost", MsgDeleteSentPost],
  ["/deep2essence.planet.blog.MsgUpdateTimedoutPost", MsgUpdateTimedoutPost],
  ["/deep2essence.planet.blog.MsgDeletePost", MsgDeletePost],
  ["/deep2essence.planet.blog.MsgDeleteTimedoutPost", MsgDeleteTimedoutPost],
  ["/deep2essence.planet.blog.MsgCreateTimedoutPost", MsgCreateTimedoutPost],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgUpdateSentPost: (data: MsgUpdateSentPost): EncodeObject => ({ typeUrl: "/deep2essence.planet.blog.MsgUpdateSentPost", value: MsgUpdateSentPost.fromPartial( data ) }),
    msgUpdatePost: (data: MsgUpdatePost): EncodeObject => ({ typeUrl: "/deep2essence.planet.blog.MsgUpdatePost", value: MsgUpdatePost.fromPartial( data ) }),
    msgCreatePost: (data: MsgCreatePost): EncodeObject => ({ typeUrl: "/deep2essence.planet.blog.MsgCreatePost", value: MsgCreatePost.fromPartial( data ) }),
    msgCreateSentPost: (data: MsgCreateSentPost): EncodeObject => ({ typeUrl: "/deep2essence.planet.blog.MsgCreateSentPost", value: MsgCreateSentPost.fromPartial( data ) }),
    msgDeleteSentPost: (data: MsgDeleteSentPost): EncodeObject => ({ typeUrl: "/deep2essence.planet.blog.MsgDeleteSentPost", value: MsgDeleteSentPost.fromPartial( data ) }),
    msgUpdateTimedoutPost: (data: MsgUpdateTimedoutPost): EncodeObject => ({ typeUrl: "/deep2essence.planet.blog.MsgUpdateTimedoutPost", value: MsgUpdateTimedoutPost.fromPartial( data ) }),
    msgDeletePost: (data: MsgDeletePost): EncodeObject => ({ typeUrl: "/deep2essence.planet.blog.MsgDeletePost", value: MsgDeletePost.fromPartial( data ) }),
    msgDeleteTimedoutPost: (data: MsgDeleteTimedoutPost): EncodeObject => ({ typeUrl: "/deep2essence.planet.blog.MsgDeleteTimedoutPost", value: MsgDeleteTimedoutPost.fromPartial( data ) }),
    msgCreateTimedoutPost: (data: MsgCreateTimedoutPost): EncodeObject => ({ typeUrl: "/deep2essence.planet.blog.MsgCreateTimedoutPost", value: MsgCreateTimedoutPost.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
