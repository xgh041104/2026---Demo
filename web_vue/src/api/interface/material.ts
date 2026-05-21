// 分页响应参数
export interface ResPage<T> {
  list: T[];
  pageNum: number;
  pageSize: number;
  total: number;
}

export interface ResMaterial<T> {
  list: T[];
  total: number;
}

// 分页请求参数
export interface ReqPage {
  pageNum: number;
  pageSize: number;
}

export namespace Material {
  export interface ReqMaterial {
    id: number;
    name: string;
    creator_name: string;
    created_at: string;
    image_url: string;
  }

  export interface upAuditMaterial {
    id: number;
    status: number;
  }

  export interface ReqUserParams extends ReqPage {
    account: string;
    password: string;
  }

  export interface UserList {
    username: string;
    gender: number;
  }
}
