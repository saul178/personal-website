import axios from "axios";

const axiosInstance = axios.create({
  // TODO: replace this with a proper domain name
  baseURL: import.meta.env.VITE_API_BASE_URL,
})

export default axiosInstance
