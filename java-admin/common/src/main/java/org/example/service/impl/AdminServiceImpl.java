package org.example.service.impl;

import cn.hutool.crypto.SecureUtil;
import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import lombok.extern.slf4j.Slf4j;
import org.example.dao.AdminDao;
import org.example.entity.Admin;
import org.example.service.AdminService;
import org.example.util.AssertUtil;
import org.example.util.HttpServletUtil;
import org.example.util.TokenUtil;
import org.springframework.stereotype.Service;

import java.util.HashMap;


@Service
@Slf4j
public class AdminServiceImpl extends ServiceImpl<AdminDao, Admin> implements AdminService {

    @Override
    public Object login(Admin v) {
        Admin admin = this.getOne(new LambdaQueryWrapper<Admin>().eq(Admin::getUserName, v.getUserName()));
        if (admin == null || !admin.getPassword().equals(getMd5Password(v.getPassword(), admin.getSalt()))) {
            AssertUtil.AssertError("密码或账号错误");
        }
        HashMap<String, Object> map = new HashMap<>();
        map.put("token", TokenUtil.getToken(admin.getId()));
        return map;
    }

    @Override
    public Object getUserInfo() {
        Admin admin = getCurrentUser();
        HashMap<String, Object> map = new HashMap<>();
        map.put("userId", admin.getId());
        map.put("username", "管理员");
        map.put("realName", "管理员");
        return map;
    }


    public static String getMd5Password(String password, String salt) {
        // md5加密
        return SecureUtil.md5(password + salt);
    }

    // If login success, then return userId
    public String getUserId() {
        String token = HttpServletUtil.getHeader(TokenUtil.Authorize);
        return TokenUtil.getUserId(token);
    }

    // Get current user
    public Admin getCurrentUser() {
        return this.getById(getUserId());
    }

}
