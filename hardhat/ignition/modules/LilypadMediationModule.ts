import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

export default buildModule("LilypadMediationRandomModule", (m) => {
  const admin = m.getAccount(0);
  const mediation = m.contract("LilypadMediationRandom", [], {
    from: admin,
  });
  m.call(mediation, "initialize", [], {
    from: admin,
  });

  return { mediation };
});