import axios, { AxiosResponse } from 'axios';

interface LinkShortenAPIBody {
  domain: string;
  uri: string;
  target: string;
  token: string;
}

const ShortenLinkAPI = async (
  body: LinkShortenAPIBody,
  onSuccess: (res: AxiosResponse) => void,
  onError: (res: AxiosResponse) => void,
) => {
  const axiosInstance = axios.create({
    baseURL: document.location.protocol + '//' + document.location.host,
  });
  const response = await axiosInstance.post('/api/shorten', body);
  if (response.status === 200) {
    onSuccess(response);
    return;
  }
  onError(response);
  return;
};

export default ShortenLinkAPI;
