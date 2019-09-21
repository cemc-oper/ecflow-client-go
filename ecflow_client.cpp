#include "ecflow_client.h"

#include <ClientInvoker.hpp>
#include <Defs.hpp>
//#include <DState.hpp>

namespace EcflowUtil {

class EcflowClientPrivate {
public:
    EcflowClientPrivate(std::string host, std::string port) :
        host_{host}, port_{port} {}

    int sync() {
        invoker_.set_host_port(host_, port_);
        auto sync_result = invoker_.sync_local();
        if (sync_result != 0) {
            error_message_ = invoker_.errorMsg();
        }
        return sync_result;
    }

    std::vector<NodeStatusRecord> collectStatus() {
        defs_ = invoker_.defs();

        std::vector<node_ptr> nodes;
        defs_->get_all_nodes(nodes);

        std::vector<NodeStatusRecord> records;

        for (auto &node: nodes) {
            NodeStatusRecord record;
            record.path_ = node->absNodePath();
            record.status_ = DState::toString(node->dstate());
            records.push_back(record);
        }

        return records;
    }

private:
    std::string host_;
    std::string port_;
    ClientInvoker invoker_;
    defs_ptr defs_;

    std::string error_message_;

    friend class EcflowClient;
};

EcflowClient::EcflowClient(const std::string &host, const std::string &port) :
    host_{host},
    port_{port},
    p_{new EcflowClientPrivate{host, port}} {

}

EcflowClient::~EcflowClient() {
    delete p_;
}

int EcflowClient::sync() {
    auto ret = p_->sync();
    if (ret != 0) {
        return ret;
    }
    status_records_ = p_->collectStatus();
    return 0;
}

std::string EcflowClient::errorMessage() {
    return p_->error_message_;
}


} // namespace EcflowUtil