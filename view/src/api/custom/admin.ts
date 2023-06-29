import { GetUserInfoModel, LoginParams } from "@/api/sys/model/userModel";
import { ErrorMessageMode } from "#/axios";
import { defHttp } from "@/utils/http/axios";

// 登入接口
export function adminLogin(params: LoginParams, mode: ErrorMessageMode = "modal") {
  return defHttp.post<any>(
    {
      url: "/admin/login",
      params
    },
    {
      errorMessageMode: mode
    }
  );
}


// 获取信息,必须需要
export function adminUserInfo() {
  return defHttp.get<GetUserInfoModel>({ url: "/admin/getUserInfo" });
}


// adminList
export function adminList(params: any) {
  return defHttp.get<any>({ url: "/admin/list", params });
}
