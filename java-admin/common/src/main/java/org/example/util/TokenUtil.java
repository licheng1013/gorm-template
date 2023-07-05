package org.example.util;

import cn.hutool.core.convert.NumberWithFormat;
import cn.hutool.core.util.StrUtil;
import cn.hutool.jwt.JWT;
import cn.hutool.jwt.JWTUtil;
import org.example.exception.ServiceException;

import java.io.Serializable;
import java.util.HashMap;
import java.util.Map;

/**
 * JWT的封装用于更易用
 * @author lc
 * @since 7/11/22
 */
public class TokenUtil  {
    public static final String Authorize = "Authorization";

    /** 密钥串 **/
    public static String SECRET_KEY = "YOUR_SECRET_KEY";
    private static final  String USER_ID_KEY = "USER_ID";
    private static final  String EXPIRE_TIME_KEY = "EXPIRE_TIME";
    /**
     * 获取Token
     * @since 7/11/22
     */
    public static String getToken(Serializable userId){
        Map<String, Object> map = new HashMap<>();
        map.put(USER_ID_KEY, userId);
        map.put(EXPIRE_TIME_KEY, System.currentTimeMillis() + 1000 * 60 * 60 * 24 * 3);
        return JWTUtil.createToken(map, SECRET_KEY.getBytes());
    }

    /**
     * 获取用户id
     * @since 7/11/22
     */
    public static String getUserId(String token){
        if (StrUtil.isBlankIfStr(token)) {
            throw new ServiceException("令牌为空!",-10);
        }

        JWT jwt = JWTUtil.parseToken(token);
        if (!JWTUtil.verify(token, SECRET_KEY.getBytes())) {
            throw new ServiceException("无效Token!",-10);
        }
        Object payload = jwt.getPayload(EXPIRE_TIME_KEY);
        if (payload == null || System.currentTimeMillis() > ((NumberWithFormat)payload).longValue()) {
            throw new ServiceException("登入失效!",-10);
        }
        return jwt.getPayload(USER_ID_KEY).toString();
    }

    public static void main(String[] args) {
        //测试
        String token = getToken(123456);
        System.out.println(token);
        String userId = getUserId(token);
        System.out.println(userId);
    }
}
