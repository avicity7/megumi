import { PUBLIC_BACKEND_URL } from '$env/static/public'

export const load = () => {
  const backend_uri = PUBLIC_BACKEND_URL

  return {
    backend_uri
  }
}