import {Node} from "@vervstack/matreshka";
import {GrpcClient} from "@/models/configs/verv/Resources/Resource.ts";
import {extractStringValue} from "@/models/shared/common.ts";

export function mapGrpc(root: Node): GrpcClient {
    if (!root.name) {
        throw {message: 'Can\'t parse grpc client config'}
    }

    const grpcClient = new  GrpcClient(root.name.slice(root.name.indexOf('GRPC')).toLowerCase())

    root.innerNodes?.map(
        (n) => {
            if (!n.name || !root.name) {
                return
            }

            const fieldName = n.name.slice(root.name.length + 1)

            switch (fieldName) {
                case 'CONNECTION-STRING':
                    grpcClient.connectionString = extractStringValue(n)
                    break
                case 'MODULE':
                    grpcClient.module = extractStringValue(n)
                    break
            }
        }
    )

    return grpcClient
}
