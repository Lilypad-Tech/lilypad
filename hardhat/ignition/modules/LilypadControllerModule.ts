import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";
import LilypadTokenModule from "./LilypadTokenModule";
import LilypadStorageModule from "./LilypadStorageModule";
import LilypadMediationModule from "./LilypadMediationModule";
import LilypadUsersModule from "./LilypadUsersModule";
import LilypadPaymentsModule from "./LilypadPaymentsModule";
import LilypadOnChainJobCreatorModule from "./LilypadOnChainJobCreatorModule";

export default buildModule("LilypadControllerModule", (m) => {
  const admin = m.getAccount(0);
  const solver = m.getAccount(1);
  const { token } = m.useModule(LilypadTokenModule);
  const { storage } = m.useModule(LilypadStorageModule);
  const { mediation } = m.useModule(LilypadMediationModule);
  const { users } = m.useModule(LilypadUsersModule);
  const { payments } = m.useModule(LilypadPaymentsModule);
  const { jobCreator } = m.useModule(LilypadOnChainJobCreatorModule);

  const controller = m.contract("LilypadController", [], {
    from: admin,
  });

  m.call(controller, "initialize", [storage, mediation, users, payments, jobCreator], {
    from: admin,
  });

  m.call(storage, "setControllerAddress", [controller], {
    from: admin,
  });

  m.call(payments, "setControllerAddress", [controller], {
    from: admin,
  });
  
  m.call(mediation, "setControllerAddress", [controller], {
    from: admin,
  });

  return { controller };
});