<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col lg="6" offset-lg="3" md="8" offset-md="2">
        <b-card title="登录">
          <b-form>
            <b-form-group label="手机号">
              <b-form-input
                v-model="$v.user.telphone.$model"
                type="number"
                required
                placeholder="请输入名称"
              ></b-form-input>
            </b-form-group>
            <b-form-valid-feedback :state="validateState('telphone')"
              >手机号不合法</b-form-valid-feedback
            >

            <b-form-group label="密码">
              <b-form-input
                v-model="$v.user.password.$model"
                type="password"
                required
                placeholder="请输入密码"
              ></b-form-input>
            </b-form-group>
            <b-form-valid-feedback :state="validateState('password')"
              >密码不能小于六个字符</b-form-valid-feedback
            >
            <b-form-group class="mt-5">
              <b-button @click="login" variant="outline-primary" block
                >登录</b-button
              >
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { required, minLength, maxLength } from "vuelidate/lib/validators";
import validator from "../../helper/validator";
import { mapActions } from "vuex";

export default {
  data() {
    return {
      user: {
        telphone: "",
        password: "",
      },
    };
  },
  validations: {
    user: {
      password: {
        required,
        minLength: minLength(6),
      },
      telphone: {
        telphone: validator.telphoneValidate,
      },
    },
  },
  methods: {
    validateState(field) {
      let { $dirty, $error } = this.$v.user[field];
      return $dirty ? $error : null;
    },
    ...mapActions("userModule", { userLogin: "login" }),
    login() {
      //验证数据
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      this.userLogin(this.user) //this.$store.dispatch("userModule/register", this.user)
        .then((res) => {
          console.log("userModule/userLogin", res);
          if (res.data.user && res.status == 200) {
            this.$router.replace({ name: "Home" });
          }
          if (res.data.code != 200) {
            this.$bvToast.toast(res.data.msg, {
              title: `温馨提示`,
              variant: "warning",
              solid: true,
            });
          }
        })
        .catch((err) => {
          this.$bvToast.toast(JSON.stringify(err), {
            title: `网络异常`,
            variant: "danger",
            solid: true,
          });
          console.log(err);
        });
    },
  },
};
</script>
