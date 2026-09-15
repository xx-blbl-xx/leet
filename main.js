const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const path = require('path');

// 加载 proto 文件
const PROTO_PATH = path.join(__dirname, 'api.proto');
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
  includeDirs: [
    // 如果有依赖的 proto 文件路径,需要添加到这里
    // 例如: '/path/to/proto/dependencies'
  ]
});

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
const operateService = protoDescriptor.customerservice.operateservice.v1;

// 创建客户端
const client = new operateService.OperatePlatformService(
  '10.150.98.228:9000', // 替换为实际的服务地址
  grpc.credentials.createInsecure() // 生产环境请使用安全凭证
);

// 构造请求参数
const request = {
  init_time: {
    gte: 1761537483, // 开始时间戳(必填)
    lte: 1764215883  // 结束时间戳(必填,限3个月内)
  },
//   business_type: [7617, 7618, 7619, 7744], // 业务类型id数组(可选)
  fans: {
    gte: 0,      // 粉丝数下限(可选)
    lte: 100000  // 粉丝数上限(可选)
  },
  from: 1,   // 分页起始位置(必填,从1开始)
  size: 20   // 每页大小(必填,最大100)
};

// 调用 gRPC 方法
client.SearchConversationFromEs(request, (error, response) => {
  if (error) {
    console.error('调用失败:', error);
    return;
  }
  
  console.log('调用成功:');
  console.log('会话列表:', response.conversation_list);
  
  // 遍历返回的会话数据
  if (response.conversation_list && response.conversation_list.length > 0) {
    response.conversation_list.forEach((conversation, index) => {
      console.log(`\n--- 会话 ${index + 1} ---`);
      console.log('会话ID:', conversation.conversation_id);
      console.log('用户名:', conversation.customer_name);
      console.log('客服ID:', conversation.staff_id);
      console.log('客服内部名称:', conversation.staff_inner_name);
      console.log('客服外显名称:', conversation.staff_display_name);
      console.log('会话时长:', conversation.session_duration);
      console.log('排队时长:', conversation.queue_time);
      console.log('创建时间:', conversation.init_time);
      console.log('结束时间:', conversation.end_time);
      console.log('评价类型:', conversation.eva_type);
      console.log('用户ID:', conversation.mid);
      console.log('粉丝数:', conversation.fans);
      
      // 会话总结
      if (conversation.summary) {
        console.log('会话总结:');
        console.log('  业务名称:', conversation.summary.business_name);
        console.log('  业务类型:', conversation.summary.business_type_name);
        console.log('  备注:', conversation.summary.remark);
        console.log('  处理状态:', conversation.summary.processing_status_str);
      }
    });
  }
});

// 使用 Promise 方式调用
function searchConversationFromEsAsync(request) {
  return new Promise((resolve, reject) => {
    client.SearchConversationFromEs(request, (error, response) => {
      if (error) {
        reject(error);
      } else {
        resolve(response);
      }
    });
  });
}

// Promise 方式使用示例
async function main() {
  try {
    const result = await searchConversationFromEsAsync(request);
    console.log('异步调用成功:', result);
  } catch (error) {
    console.error('异步调用失败:', error);
  }
}

// 取消注释以运行异步示例
main();