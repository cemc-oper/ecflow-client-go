#include "ecflow_util.h"

#include <ClientInvoker.hpp>
#include <Defs.hpp>
//#include <DState.hpp>

namespace EcflowUtil {

class EcflowClientWrapperPrivate {
public:
    EcflowClientWrapperPrivate(std::string host, std::string port) :
        host_{host}, port_{port} {
        invoker_.set_throw_on_error(false);
    }

    void setConnectTimeout(int time_out) {
        invoker_.set_connect_timeout(time_out);
    }

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

        NodeStatusRecord record;
        record.path_ = "/";
        record.status_ = NState::toString(defs_->state());
        records.push_back(record);

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

    friend class EcflowClientWrapper;
};

EcflowClientWrapper::EcflowClientWrapper(const std::string &host, const std::string &port) :
    host_{host},
    port_{port},
    p_{new EcflowClientWrapperPrivate{host, port}} {

}

EcflowClientWrapper::~EcflowClientWrapper() {
    delete p_;
}

void EcflowClientWrapper::setConnectTimeout(int time_out) {
    p_->setConnectTimeout(time_out);
}

int EcflowClientWrapper::sync() {
    auto ret = p_->sync();
    if (ret != 0) {
        return ret;
    }
    status_records_ = p_->collectStatus();
    return 0;
}

std::string EcflowClientWrapper::errorMessage() {
    return p_->error_message_;
}


} // namespace EcflowUtil