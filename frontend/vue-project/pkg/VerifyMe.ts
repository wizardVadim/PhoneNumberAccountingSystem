import { request } from './Request'
import router from '../src/router'

export const verify = async (): Promise<boolean> => {
  const token = localStorage.getItem('token')
  
  if (!token) {
    router.push('/login')
    return false
  }
  
  try {
    const response = await request('/auth/verify', { method: 'GET' })
    console.log("Token valid", response)
    
    if (response.valid === true) {
      localStorage.setItem('roleId', response.role_id)
      console.log("response.role_id", response.role_id)
      return true
    } else {
      localStorage.removeItem('token')
      localStorage.removeItem('roleId')
      router.push('/login')
      return false
    }
  } catch (error) {
    console.error("Verify error:", error)
    localStorage.removeItem('token')
    localStorage.removeItem('roleId')
    router.push('/login')
    return false
  }
}