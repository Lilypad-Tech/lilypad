import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";
import LilypadTokenModule from "./LilypadTokenModule";

export default buildModule("LilypadUsersModule", (m) => {
  const admin = m.getAccount(0);
  const users = m.contract("LilypadUsers", [], {
    from: admin,
  });
  m.call(users, "initialize", [], {
    from: admin,
  });

  return { users };
});