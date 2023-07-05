package org.example.controller;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import org.example.entity.Admin;
import org.example.service.AdminService;
import org.example.util.R;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;


@RestController
@RequestMapping("/admin")
public class AdminController {

    @Autowired
    private AdminService vService;


    @GetMapping("/list") //列表
    public Object list(Admin v){
        return R.okData(vService.list(new QueryWrapper<>(v)));
    }
    @GetMapping("/one") //单条记录
    public Object one(String id){
        return R.okData(vService.getById(id));
    }
    @GetMapping("/page") //分页
    public Object page(Page<Admin> page,Admin v){
        return R.okData(vService.page(page, new QueryWrapper<>(v)));
    }
    @PostMapping("/delete") //多条删除 => 1,2,3
    public Object deleteAll(@RequestBody ArrayList<String> ids){
        vService.removeByIds(ids);
        return R.okMsg("删除所有成功!");
    }
    @PostMapping("/insert") //插入
    public Object insert(@RequestBody Admin v){
        vService.save(v);
        return R.okMsg("插入成功!");
    }
    @PostMapping("/update") //修改
    public Object update(@RequestBody Admin v){
        vService.updateById(v);
        return R.okMsg("修改成功!");
    }
}
