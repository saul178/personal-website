import axios from "axios";

const axiosInstance = axios.create({
  // NOTE: change this back to /api when in prod, try to find a better way to switch between dev and
  // prod
  baseURL: "http://localhost:8080/api",
})

export default axiosInstance
