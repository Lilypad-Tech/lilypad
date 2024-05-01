import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";
import LilypadTokenModule from "./LilypadTokenModule";

export default buildModule("LilypadOnChainJobCreatorModule", (m) => {
  const admin = m.getAccount(0);
  const solver = m.getAccount(1);
  const { token } = m.useModule(LilypadTokenModule);
  const jobCreator = m.contract("LilypadOnChainJobCreator", [], {
    from: admin,
  });
  
  const example = m.contract("ExampleClient", [], {
    from: admin,
  });

  m.call(jobCreator, "initialize", [token], {
    from: admin,
  });

  m.call(example, "initialize", [jobCreator], {
    from: admin,
  });

  m.call(jobCreator, "setControllerAddress", [solver], {
    from: admin,
  });

  return { jobCreator, example };
});