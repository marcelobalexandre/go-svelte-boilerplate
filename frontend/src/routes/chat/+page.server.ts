export const load = ({ cookies }) => {
  return {
    token: cookies.get('token')
  }
}
