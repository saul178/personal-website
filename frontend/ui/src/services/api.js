import axios from "axios";

const axiosInstance = axios.create({
  // NOTE: change this back to /api when in prod, try to find a better way to switch between dev and
  // prod
  baseURL: "/api",
})

export default axiosInstance
