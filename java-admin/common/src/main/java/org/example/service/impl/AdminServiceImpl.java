package org.example.service.impl;

import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.example.dao.AdminDao;
import org.example.entity.Admin;
import org.example.service.AdminService;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;


@Service
@Transactional
public class AdminServiceImpl extends ServiceImpl<AdminDao, Admin> implements AdminService {

}
