package org.example.config;

import com.baomidou.mybatisplus.core.handlers.MetaObjectHandler;
import lombok.extern.slf4j.Slf4j;
import org.apache.ibatis.reflection.MetaObject;
import org.springframework.stereotype.Component;

import java.time.LocalDateTime;
import java.util.Date;

/**
 * 参考: <a href="https://baomidou.com/pages/4c6bcf/">...</a>
 * @author lc
 * @since 2023/1/14
 */
@Component
@Slf4j
public class MyMetaObjectHandler implements MetaObjectHandler {

    @Override
    public void insertFill(MetaObject metaObject) {
        log.info("插入执行 ....");
        this.setFieldValByName("createTime", new Date(), metaObject);
    }

    @Override
    public void updateFill(MetaObject metaObject) {
        log.info("更新执行 ....");
        this.setFieldValByName("updateTime", new Date(), metaObject);
    }

}