import type { X509Certificate } from 'crypto'

export default (cert: X509Certificate) => {
  const formatDN = (dn: string) => {
    return dn.split('\n').join(',')
  }

  return `x509::${formatDN(cert.subject)}::${formatDN(cert.issuer)}`
}
