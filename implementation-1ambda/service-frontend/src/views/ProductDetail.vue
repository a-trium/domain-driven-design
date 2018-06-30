<template>
    <div>
        <el-row type="flex" justify="center">
            <el-col :xs="20" :sm="20" :md="18" :lg="18">
                <div>
                    <h1>Product Detail</h1>
                </div>
            </el-col>
        </el-row>
        <el-row type="flex" justify="center" class="detail">
        </el-row>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapActions, mapGetters, mapMutations, mapState } from 'vuex'
import { Exception, Product as ProductItem, ProductOption,} from '@/generated/swagger'
import { ProductAPI, } from '@/common/product.service.ts'

@Component({
    components: {},
    computed: {
        ...mapState([ 'uid' ]),
        ...mapGetters([ 'authenticated' ]),
    },
})
export default class ProductDetail extends Vue {
    public $refs: any
    public $notify: any
    public $router: any
    public $route: any
    public $store: any

    public productID: string = ""
    public productItem: ProductItem = null
    public productOptions: Array<ProductOption> = []

    mounted() {
        this.productID = this.$route.params.productID
        ProductAPI.findOneWithOptions(this.productID, { credentials: 'include' })
            .then((response: any) => {
                this.productItem = response.product
                this.productOptions = response.options
            })
            .catch((response) => {
                if (!response.json) {
                    this.$notify.error({
                        title: `Error (Connection)`,
                        message: "Server is not available",
                    })

                    return
                }

                response.json().then((parsed: Exception) => {
                    this.$notify.error({
                        title: `Error (${parsed.type})`,
                        message: parsed.message,
                    })
                })
            })
    }
}
</script>

<style>
.detail {
    margin-top: 30px;
    text-align: center;
}
</style>