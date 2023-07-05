package org.example.controller;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import org.example.entity.Admin;
import org.example.middleware.PassToken;
import org.example.service.AdminService;
import org.example.util.AssertUtil;
import org.example.util.PageDto;
import org.example.util.PageResult;
import org.example.util.R;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;


@RestController
@RequestMapping("/admin")
public class AdminController {

    @Autowired
    private AdminService vService;


    @GetMapping("/list") //分页与go语言同步
    public Object page(PageDto page, Admin v) {
        Page<Admin> adminPage = vService.page(new Page<>(page.getPage(), page.getSize()), new QueryWrapper<>(v));
        return R.okData(new PageResult(adminPage.getRecords(), adminPage.getTotal()));
    }

    @PostMapping("/delete") //多条删除 => 1,2,3
    public Object deleteAll(@RequestBody PageDto page) {
        AssertUtil.AssertError("请从后端取消删除");
        vService.removeByIds(page.getIds());
        return R.okMsg("删除所有成功!");
    }

    @PostMapping("/insert") //插入
    public Object insert(@RequestBody Admin v) {
        vService.save(v);
        return R.okMsg("插入成功!");
    }

    @PostMapping("/update") //修改
    public Object update(@RequestBody Admin v) {
        vService.updateById(v);
        return R.okMsg("修改成功!");
    }

    // login
    @PostMapping("/login")
    @PassToken
    public Object login(@RequestBody Admin v) {
        return R.okData(vService.login(v));
    }

    // getUserInfo
    @GetMapping("/getUserInfo")
    public Object getUserInfo() {
        return R.okData(vService.getUserInfo());
    }

}
