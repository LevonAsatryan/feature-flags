import { Injectable } from '@nestjs/common';
import { CreateContainerDto } from './dto/create-container.dto';
import { UpdateContainerDto } from './dto/update-container.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { Container } from './entities/container.entity';
import { Repository } from 'typeorm';
import { ReturnContainerDto } from './dto/return-container.dto';

@Injectable()
export class ContainersService {
  constructor(
    @InjectRepository(Container)
    private containerRepository: Repository<Container>
  ) {}

  async create(createContainerDto: CreateContainerDto) {
    const container = new Container();
    container.name = createContainerDto.name;
    container.companyId = createContainerDto.companyId;

    const parentContainer = await this.containerRepository.find({
      where: { id: createContainerDto.parentContainerId },
    });
    container.parentContainer = parentContainer[0];

    return this.containerRepository.save(container);
  }

  createRootContainer(companyId: number) {
    const container = new Container();
    container.name = 'root';
    container.companyId = companyId;
    container.parentContainer = null;
    return this.containerRepository.save(container);
  }

  findAll() {
    return this.containerRepository.find({ relations: ['parentContainer'] });
  }

  async findOne(id: number) {
    const returnContainer = new ReturnContainerDto();
    const container = await this.containerRepository.findOne({
      where: { id },
      relations: ['parentContainer'],
    });

    returnContainer.id = container.id;
    returnContainer.name = container.name;
    returnContainer.parentContainer = container.parentContainer;

    const childrenContainers = (await this.containerRepository.find({
      where: { parentContainer: { id } },
      loadRelationIds: true,
    })) as ReturnContainerDto[];

    returnContainer.childrenContainers = childrenContainers;
    return returnContainer;
  }

  async update(id: number, updateContainerDto: UpdateContainerDto) {
    const container = await this.containerRepository.findOne({
      where: { id },
      relations: ['parentContainer'],
    });
    container.name = updateContainerDto.name || container.name;

    const parentContainer = await this.containerRepository.findOne({
      where: { parentContainer: { id: updateContainerDto.parentContainerId } },
    });

    container.parentContainer = parentContainer || container.parentContainer;
    return this.containerRepository.save(container);
  }

  remove(id: number) {
    return this.containerRepository.delete({ id });
  }
}
