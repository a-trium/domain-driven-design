<template>
    <div>
        <el-row type="flex" justify="center">
            <el-col :xs="20" :sm="20" :md="18" :lg="18">
                <span style="margin-left: 8px; margin-right: 8px;">
                    <router-link :to="{name: 'product'}">
                    <el-button icon="el-icon-menu" :to="'/product'">
                        Products
                    </el-button>
                    </router-link>
                </span>
                <template v-for="(category, index) in currentCategories">
                    <span style="margin-left: 8px; margin-right: 8px;">
                        <i class="el-icon-arrow-right"></i>
                    </span>
                    <el-select v-model="currentCategories[index]"
                               placeholder="Select">
                        <el-option
                                :key="optionsPerCategory[index][0].value"
                                :label="optionsPerCategory[index][0].label"
                                :value="optionsPerCategory[index][0].value">
                        </el-option>
                    </el-select>
                </template>
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

    public currentCategories = []
    public optionsPerCategory = []

    mounted() {
        this.productID = this.$route.params.productID
        ProductAPI.findOneWithOptions(this.productID, { credentials: 'include' })
            .then((response: any) => {
                this.productItem = response.product
                this.productOptions = response.options
                const filtered = response.product.categoryPath.split("/").filter(x => x.trim() !== "")
                this.currentCategories = filtered

                // TODO: List all available categories for navigation
                this.optionsPerCategory = filtered.map(x => [{ label: x, value: x }])

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