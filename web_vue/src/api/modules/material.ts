import { ReqPage, Material, ResMaterial } from "@/api/interface/material";
import http from "@/api";

export const getMaterialList = (params: ReqPage) => {

  return http.put<ResMaterial<Material.ReqMaterial>>(`/v1/UpdateAuditingMaterialList`, params);
};

export const upAuditMaterialApi = (params: { id: number; status: number; remark: string }) => {
  return http.post(`/v1/UpAuditMaterial`, params);
};

export const saveMaterialApi = (params: FormData) => {
  return http.post(`/v1/SaveMaterial`, params);
};