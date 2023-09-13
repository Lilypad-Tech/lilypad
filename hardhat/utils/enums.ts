// IMPORTANT: keep these lined up with hardhat/contracts/SharedStructs.sol
// TODO: automate this
export const ServiceType = [
  'Solver',
  'Mediator',
  'ResourceProvider',
  'JobCreator',
]

export const AgreementState = [
  'DealNegotiating',
  'DealAgreed',
  'ResultsSubmitted',
  'ResultsAccepted',
  'ResultsChecked',
  'MediationAccepted',
  'MediationRejected',
  'TimeoutAgree',
  'TimeoutSubmitResults',
  'TimeoutJudgeResults',
  'TimeoutMediateResults',
]


export const PaymentReason = [
  'PaymentCollateral',
  'ResultsCollateral',
  'TimeoutCollateral',
  'JobPayment',
  'MediationFee',
]

export const PaymentDirection = [
  'PaidIn',
  'PaidOut',
  'Refunded',
  'Slashed',
]

export const getTypeIndex = (name: string, arr: string[], type: string) => {
  const ret = arr.indexOf(type)
  if(ret < 0) throw new Error(`no ${name} of type ${type}`)
  return ret
}
export const getServiceType = (type: string) => getTypeIndex('ServiceType', ServiceType, type)
export const getAgreementState = (type: string) => getTypeIndex('AgreementState', AgreementState, type)
export const getPaymentReason = (type: string) => getTypeIndex('PaymentReason', PaymentReason, type)
export const getPaymentDirection = (type: string) => getTypeIndex('PaymentDirection', PaymentDirection, type)