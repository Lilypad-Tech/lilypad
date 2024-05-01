import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";
import LilypadTokenModule from "./LilypadTokenModule";

export default buildModule("LilypadPaymentsModule", (m) => {
  const admin = m.getAccount(0);
  const { token } = m.useModule(LilypadTokenModule);
  const payments = m.contract("LilypadPayments", [], {
    from: admin,
  });
  const _paymentsCall = m.call(payments, "initialize", [token], {
    from: admin,
  });

  const _executeCall = m.call(token, "setControllerAddress", [payments], {
    from: admin,
  });

  return { payments };
});