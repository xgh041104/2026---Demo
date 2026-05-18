import { ReqPage, Material, ResMaterial } from "@/api/interface/material";
import http from "@/api";

export const getMaterialList = (params: ReqPage) => {

  return http.get<ResMaterial<Material.ReqMaterial>>(`/GetAuditingMaterialList`, params);
};

export const upAuditMaterialApi = (params: { id: number; status: number; remark: string }) => {
  return http.post(`/v1/UpAuditMaterial`, params);
};
