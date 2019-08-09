import grpc
from concurrent import futures
import time
import math
from models import calculator_pb2
from models import calculator_pb2_grpc

class Calculator(calculator_pb2_grpc.CalculatorServicer):
    def SquareRoot(self, request ,context):
        response = calculator_pb2.Number()
        response.value =  math.sqrt(request.value)
        return response
server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
calculator_pb2_grpc.add_CalculatorServicer_to_server(
    Calculator(), server
    )
print('Starting server. Listening on port 50051')
server.add_insecure_port('localhost:50051')
server.start()
try:
    while True:
        time.sleep(86400)
except KeyboardInterrupt:
    server.stop(0)