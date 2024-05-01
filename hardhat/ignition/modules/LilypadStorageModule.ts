import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

export default buildModule("LilypadStorageModule", (m) => {
  const admin = m.getAccount(0);
  const storage = m.contract("LilypadStorage", [], {
    from: admin,
  });
  const storageCall = m.call(storage, "initialize", [], {
    from: admin,
  });

  return { storage };
});