package org.example.entity;

import lombok.Data;
import com.baomidou.mybatisplus.annotation.*;
import java.util.Date;


@Data
public class Admin {

    /** 管理员id */
    @TableId(type = IdType.AUTO)
    private Long id;
    /** 账号 */
    private String userName;
    /** 密码 */
    private String password;
    /** 盐 */
    private String salt;
    /** 创建时间 */
    @TableField(fill = FieldFill.INSERT)
    private Date createTime;
    /** 昵称 */
    private String nickName;
    
}