import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";
import {
  DEFAULT_TOKEN_SUPPLY,
} from '../../utils/web3'

export default buildModule("LilypadTokenModule", (m) => {
  const admin = m.getAccount(0);
  const token = m.contract("LilypadToken", ["Lilypad Token", "LP", DEFAULT_TOKEN_SUPPLY], {
    from: admin,
  });
  return { token };
});