import bluebird from 'bluebird'
import {
  connectUsers,
} from '../utils/web3'

import {
  getServiceType,
} from '../utils/enums'

async function main() {
  const storage = await connectUsers()

  const results = await bluebird.props({
    Solver: storage.showUsersInList(getServiceType('Solver')),
    Mediator: storage.showUsersInList(getServiceType('Mediator')),
    ResourceProvider: storage.showUsersInList(getServiceType('ResourceProvider')),
    JobCreator: storage.showUsersInList(getServiceType('JobCreator')),
  })

  console.log(JSON.stringify(results, null, 4))
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
