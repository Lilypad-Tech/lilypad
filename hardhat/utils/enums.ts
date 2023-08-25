// IMPORTANT: keep these lined up with hardhat/contracts/SharedStructs.sol
// TODO: automate this
export const ServiceType = [
  'Solver',
  'Mediator',
  'Directory',
  'ResourceProvider',
  'JobCreator'
]

export const AgreementState = [
  'Negotiating',
  'Agreed',
  'Timeout',
  'Submitted',
  'Accepted',
  'Rejected'
]

export const PaymentDirection = [
  'In',
  'Out',
  'Slashed'
]

export const PaymentReason = [
  'JobCollateral',
  'JobCollateralRefund',
  'JobCollateralTimeoutRefund',
  'ResultsCollateral',
  'ResultsCollateralRefund',
  'ResultsCollateralSlashed',
  'JobPayment',
  'TimeoutCollateral',
  'TimeoutCollateralRefund',
  'TimeoutCollateralSlashed'
]

export const getServiceType = (type: string) => ServiceType.indexOf(type)
export const getAgreementState = (type: string) => AgreementState.indexOf(type)
export const getPaymentDirection = (type: string) => PaymentDirection.indexOf(type)
export const getPaymentReason = (type: string) => PaymentReason.indexOf(type)