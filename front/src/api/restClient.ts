import axios, { AxiosResponse } from 'axios';
import qs from 'qs';

const API_URL = 'http://localhost:3000/api';

export interface IErrorHandler {
  errorCode: number;
  handler: (res: AxiosResponse) => IRequestError;
}

export interface IResponse {
  message: string;
  data: object;
}

export interface IRequestError {
  raisedAt: 'client' | 'server' | 'unknown';
  code: number;
  message: string;
  detail: string;
}

export default class BERestClient {
  apiClient;
  responseHandler;
  errorHandlers: IErrorHandler[] = [];

  constructor(
    resHandler: (res: AxiosResponse) => IResponse,
    accessToken?: string,
  ) {
    this.responseHandler = resHandler;
    this.apiClient = axios.create({
      baseURL: API_URL,
      withCredentials: false,
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
        Authorization: `Bearer ${accessToken}`,
      },
    });
  }

  setErrorHandler(
    errorCode: number,
    handler: (res: AxiosResponse) => IRequestError,
  ) {
    this.errorHandlers.push({ errorCode: errorCode, handler: handler });
  }

  runErrorHandler(errorCode: number, res: AxiosResponse): IRequestError {
    const handler = this.errorHandlers.find(
      (handler) => handler.errorCode === errorCode,
    );
    if (handler) {
      return handler.handler(res);
    } else {
      return {
        raisedAt: 'unknown',
        code: 0,
        message: 'Unknown Error',
        detail: 'Unknown Error',
      };
    }
  }

  processResponse(res: AxiosResponse): IResponse | IRequestError {
    if (res.status === 200) {
      return this.responseHandler(res);
    } else {
      return this.runErrorHandler(res.status, res);
    }
  }

  async get(path: string, params: object, qsOption?: qs.IStringifyOptions) {
    const res = await this.apiClient.get(path, {
      params: qs.stringify(params, qsOption),
    });
    return this.processResponse(res);
  }
  async post(path: string, body: object) {
    const res = await this.apiClient.post(path, body);
    return this.processResponse(res);
  }
  async delete(path: string, params: object, qsOption?: qs.IStringifyOptions) {
    const res = await this.apiClient.delete(path, {
      params: qs.stringify(params, qsOption),
    });
    return this.processResponse(res);
  }
}
